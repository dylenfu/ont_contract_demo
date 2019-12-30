OntCversion = '2.0.0'
"""
dApp ont id bind wallet smart contract
"""

from ontology.builtins import concat, state
from ontology.interop.Ontology.Native import Invoke
from ontology.interop.Ontology.Contract import Migrate
from ontology.interop.Ontology.Runtime import Base58ToAddress, AddressToBase58
from ontology.interop.System.Runtime import CheckWitness, Notify, Serialize
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
KEY_ONT_ID = "k2_"
KEY_ACCOUNT = "k3_"
KEY_STAT = "k4"
VERSION = 0
INVOKE_SUCCESS = b'\x01'


def Main(op, args):
    if op == "init":
        return init()

    elif op == "bind":
        ont_id = args[0]
        account = args[1]
        return bind(ont_id, account)

    elif op == "unbind":
        ont_id = args[0]
        account = args[1]
        return unbind(ont_id, account)

    elif op == "get_ont_id":
        account = args[0]
        return get_ont_id(account)

    elif op == "get_account":
        ont_id = args[0]
        return get_account(ont_id)

    elif op == "get_owner":
        return get_owner()

    elif op == "get_count":
        return get_count()

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


def bind(ont_id, account):
    """
    bind ontID with address
    :param ont_id:
    :param account: ontid owner wallet address
    :return:
    """

    assert CheckWitness(account)

    bound_ont_id = Get(ctx, concat(KEY_ACCOUNT, account))
    if bound_ont_id == ont_id:
        raise Exception("account bind to the same ont id")

    Put(ctx, concat(KEY_ACCOUNT, account), ont_id)
    Put(ctx, concat(KEY_ONT_ID, ont_id), account)

    stat(1)
    Notify(["bind", ont_id, account])

    return True


def unbind(ont_id, account):
    """
    unbind ont id with address
    :param ont_id:
    :param account: ont id owner wallet address
    :return:
    """

    assert CheckWitness(account)

    bound_ont_id = Get(ctx, concat(KEY_ACCOUNT, account))
    if not bound_ont_id:
        raise Exception("account not bind with any ont id")

    assert bound_ont_id == ont_id

    Delete(ctx, concat(KEY_ACCOUNT, account))
    Delete(ctx, concat(KEY_ONT_ID, ont_id))

    stat(-1)
    Notify(["unbind", ont_id, account])

    return True


def get_ont_id(account):
    """
    get ont id with owner address
    :param account:
    :return:
    """

    bound_ont_id = Get(ctx, concat(KEY_ACCOUNT, account))
    if not bound_ont_id:
        raise Exception("account not bind with any ont id")

    return bound_ont_id


def get_account(ont_id):
    """
    get owner address with ont id
    :param ont_id:
    :return:
    """

    bound_account = Get(ctx, concat(KEY_ONT_ID, ont_id))
    if not bound_account:
        raise Exception("can not find any account bind with ont id")

    return bound_account


def get_owner():
    """
    get admin address
    :return:
    """
    return Get(ctx, KEY_OWNER)


def get_count():
    """
    show stat wallet address number
    :return:
    """

    count = Get(ctx, KEY_STAT)
    if not count:
        count = 0

    return count


def reset(ont_id, account):
    """
    bind account with ont id, only be invoked by the admin
    :param ont_id:
    :param account:
    :return:
    """

    assert is_address(account)
    assert CheckWitness(get_owner())

    bound_ont_id = Get(ctx, concat(KEY_ACCOUNT, account))
    assert bound_ont_id != ont_id

    Put(ctx, concat(KEY_ACCOUNT, account), ont_id)
    Put(ctx, concat(KEY_ONT_ID, ont_id), account)

    Notify(["reset", ont_id, account])

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


# inner function:
def stat(inc):
    """
    wallet count number increase or decrease
    :return:
    """

    count = Get(ctx, KEY_STAT)
    if not count:
        count = 0

    count += inc

    if count < 0:
        count = 0

    Put(ctx, KEY_STAT, count)

    return True


def is_address(account):
    """
    validate account
    :param account:
    :return:
    """

    assert (len(account) == 20 and account != ZERO_ADDRESS)

    return True
