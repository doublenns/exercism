EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2  # prep time in minutes per lasagna layer


def bake_time_remaining(elasped_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - elasped_bake_time


def preparation_time_in_minutes(number_of_layers):
    """Calculate the preparation time for a lasanga.

    :param number_of_layers: int - number of layers the lasagna has.
    :return: int - preparation time (in minutes) derived from 'PREPARATION_TIME'.

    Function that takes the number of layers a lasagna has as an argument and
    returns how many minutes the lasagna needs to be prepared
    based on the `PREPARATION_TIME` each individual layer takes
    """

    return PREPARATION_TIME * number_of_layers


def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the bake time elasped.

    :param elapsed_bake_time: int - baking time already elapsed.
    :param number_of_layers: int - number of layers the lasagna has.
    :return: int - elasped bake time (in minutes) derived from
    'preparation_time_in_minutes()'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
