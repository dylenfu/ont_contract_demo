OntCversion = '2.0.0'
"""
dApp daimler storage bucket id and did mapping smart contract
"""

from ontology.builtins import concat, state
from ontology.interop.Ontology.Native import Invoke
from ontology.interop.Ontology.Contract import Migrate
from ontology.interop.Ontology.Runtime import Base58ToAddress
from ontology.interop.System.Runtime import CheckWitness, Notify, Serialize, Deserialize
from ontology.interop.System.ExecutionEngine import GetExecutingScriptHash
from ontology.interop.System.Storage import GetContext, Put, Get, Delete

ZERO_ADDRESS = bytearray(b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00')
ONT_ADDRESS = bytearray(b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01')
ONG_ADDRESS = bytearray(b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02')

# TODO(fukun): modify admin address
ADMIN = Base58ToAddress("AGoC1LptN4uRFrymW7p7Tu3dFR7Yq2RB2P")
CONTRACT_ADDRESS = GetExecutingScriptHash()
ctx = GetContext()


KEY_OWNER = "k1"
KEY_ONT = "k2_"
KEY_BUCKET = "k3_"
VERSION = 0
INVOKE_SUCCESS = b'\x01'


def Main(op, args):
    if op == "init":
        return init()

    elif op == "bind":
        ont_id = args[0]
        bucket = args[1]
        return bind(ont_id, bucket)

    elif op == "unbind":
        ont_id = args[0]
        bucket = args[1]
        return unbind(ont_id, bucket)

    elif op == "get_bucket":
        ont_id = args[0]
        return get_bucket(ont_id)

    elif op == "get_owner":
        return get_owner()

    elif op == "reset":
        ont_id = args[0]
        admin = args[1]
        return reset(ont_id, admin)

    elif op == "upgrade":
        code = args[0]
        return upgrade(code)

    return False


def init():
    """
    save admin address
    :return:
    """

    assert (CheckWitness(ADMIN))
    Put(ctx, KEY_OWNER, ADMIN)

    Notify(["init", "success"])
    return True


def bind(ont_id, bucket):
    """
    bind ontID with address
    :param ont_id:
    :param bucket: storage server bucket id or url
    :return:
    """

    assert CheckWitness(ADMIN)

    bound_bucket = Get(ctx, concat(KEY_ONT, ont_id))
    if bound_bucket == bucket:
        raise Exception("ont id bind to the same bucket")

    if not bound_bucket:
        bind_map = get_bind_map(bucket)
        bind_map[ont_id] = True
        Put(ctx, concat(KEY_ONT, ont_id), bucket)
        Put(ctx, concat(KEY_BUCKET, bucket), Serialize(bind_map))
    else:
        bound_data = Get(ctx, concat(KEY_BUCKET, bound_bucket))
        bound_map = Deserialize(bound_data)
        bound_map.remove(ont_id)
        Put(ctx, concat(KEY_BUCKET, bound_bucket), Serialize(bound_map))

        bind_map = get_bind_map(bucket)
        bind_map[ont_id] = True
        Put(ctx, concat(KEY_ONT, ont_id), bucket)
        Put(ctx, concat(KEY_BUCKET, bucket), Serialize(bind_map))

    Notify(["bind", ont_id, bucket])

    return True


def unbind(ont_id, bucket):
    """
    unbind ont id with address
    :param ont_id:
    :param bucket: bucket id or url
    :return:
    """

    assert CheckWitness(ADMIN)

    bound_bucket = Get(ctx, concat(KEY_ONT, ont_id))
    if not bound_bucket:
        raise Exception("ont id bind with nothing")

    assert bound_bucket == bucket

    bind_data = Get(ctx, concat(KEY_BUCKET, bucket))
    bind_map = Deserialize(bind_data)
    bind_map.remove(ont_id)

    Put(ctx, concat(KEY_BUCKET, bucket), Serialize(bind_map))
    Delete(ctx, concat(KEY_ONT, ont_id))

    Notify(["unbind", ont_id, bucket])

    return True


def get_bucket(ont_id):
    """
    get owner address with ont id
    :param ont_id:
    :return:
    """

    return Get(ctx, concat(KEY_ONT, ont_id))


def get_owner():
    """
    get admin address
    :return:
    """
    return Get(ctx, KEY_OWNER)


def reset(ont_id, bucket):
    """
    bind bucket with ont id, only be invoked by the admin
    :param ont_id:
    :param bucket:
    :return:
    """

    assert is_address(bucket)
    assert CheckWitness(get_owner())

    bound_bucket = Get(ctx, concat(KEY_BUCKET, bucket))
    assert bound_bucket != bucket

    if bound_bucket:
        bound_map = get_bind_map(bound_bucket)
        bound_map.remove(ont_id)
        Put(ctx, concat(KEY_BUCKET, bound_bucket), Serialize(bound_map))

    bind_map = get_bind_map(bucket)
    bind_map[ont_id] = True

    Put(ctx, concat(KEY_BUCKET, bucket), Serialize(bind_map))
    Put(ctx, concat(KEY_ONT, ont_id), bucket)

    Notify(["reset", ont_id, bucket])

    return True


def upgrade(code):
    """
    migrate data to new contract, transfer asset to admin address
    :param code:
    :return:
    """

    assert CheckWitness(ADMIN)

    # transfer ont
    ont_balance = Invoke(VERSION, ONT_ADDRESS, "balanceOf", state(CONTRACT_ADDRESS))
    ont_response = Invoke(VERSION, ONT_ADDRESS, "transfer", [state(CONTRACT_ADDRESS, ADMIN, ont_balance)])
    if ont_response != INVOKE_SUCCESS:
        raise Exception("contract upgrade - transfer ont failed")

    # transfer ong
    ong_balance = Invoke(VERSION, ONG_ADDRESS, "balanceOf", state(CONTRACT_ADDRESS))
    ong_response = Invoke(VERSION, ONG_ADDRESS, "transfer", [state(CONTRACT_ADDRESS, ADMIN, ong_balance)])
    if ong_response != INVOKE_SUCCESS:
        raise Exception("contract upgrade - transfer ong failed")

    # migrate avm
    migrate_response = Migrate(code, "", "", "", "", "", "")
    if not migrate_response:
        raise Exception("contract upgrade - migrate avm failed")

    Notify(["upgrade", "success"])

    return True

# inner function


def get_bind_map(bucket):
    bind_data = Get(ctx, concat(KEY_BUCKET, bucket))
    bind_map = None

    if not bind_data:
        bind_map = {}
    else:
        bind_map = Deserialize(bind_data)

    return bind_map

def is_address(account):
    """
    validate account
    :param account:
    :return:
    """

    assert (len(account) == 20 and account != ZERO_ADDRESS)

    return True
