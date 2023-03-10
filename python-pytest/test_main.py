from main import sum, f
import pytest

@pytest.fixture
def sum_pack() -> list[tuple[int, int]]:
    return [
        (0,1),
        (2,3),
        (5,8),
    ]

def test_main():
    assert sum(1, 2) == 3

def test_error():
    assert sum(1, 2) == 3

def test_sum_with_pack(sum_pack: list[tuple[int, int]]):
    assert sum(*sum_pack[0]) == 1
    assert sum(*sum_pack[1]) == 5
    assert sum(*sum_pack[2]) == 13

@pytest.mark.parametrize(['x', 'y', 'result'], [
    (4, 5, 9), 
    (6, 8, 14), 
    (4, 9, 13)
])
def test_a_lot_of_sums(x: int, y: int, result: int):
    assert sum(x, y) == result

@pytest.mark.parametrize(['x'], [
    (4,),
    (10,),
    (2,),
    (8,)
])
@pytest.mark.parametrize(['y'], [
    (6,), 
    (8,), 
    (10,)
])
def test_so_many_sums(x: int, y: int):
    assert (x + y) % 2 == 0

def test_mytest():
    with pytest.raises(SystemExit):
        f()

@pytest.fixture(params=[
    7,
    3,
    0,
    6,
    8,
])
def y_item(request) -> int:
    return request.param

@pytest.fixture(params=[
    5,
    3,
    2,
    1,
    9,
    13
])
def x_item(request) -> int:
    return request.param


@pytest.fixture
def redundant_sum_pack(x_item: int, y_item: int) -> tuple[int, int, int]:
    return (x_item, y_item, x_item + y_item)

def test_that_sum_sums(redundant_sum_pack: tuple[int, int, int]):
    x, y, result = redundant_sum_pack
    assert sum(x, y) == result