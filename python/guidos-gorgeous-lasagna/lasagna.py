"""
this module is related to preparation time of a lasagna
"""

EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2


def bake_time_remaining(elapsed_bake_time):
    '''
    :param elapsed_bake_time: int baking time already elapsed
    :return: int remaining bake time derived from 'EXPECTED_BAKE_TIME'

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    '''
    return EXPECTED_BAKE_TIME - elapsed_bake_time


def preparation_time_in_minutes(number_of_layers):
    '''
    :param number_of_layers: int number of layers want to add to the lasagna
    :return: int preparation time derived from 'PREPARATION_TIME'

    Function that takes the number of layers want to add to lasagna as an
    argument and returns how many minutes how many minutes would spend making
    them based on the `PREPARATION_TIME`.
    '''
    return number_of_layers * PREPARATION_TIME


def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    '''
    :param number_of_layers: int number of layers want to add to the lasagna
    :param elapsed_bake_time: int baking time already elapsed
    :return: int elapsed cooking time

    Function that has two parameters: number_of_layers (the number of layers 
    added to the lasagna) and elapsed_bake_time (the number of minutes the
    lasagna has been baking in the oven). This function returns the total
    number of minutes the lasagna has been cooking (the sum of preparation time
    and the time the lasagna has already spent backing in the oven.)
    '''
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
