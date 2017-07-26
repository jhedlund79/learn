//
// Created by Jason Hedlund on 5/19/17.
//

#include "celcius_to_farenheit_test.h"
#include <libunittest/all.hpp>
#include "../src/temp_conversion.hpp"

using namespace unittest::assertion;

TEST(test_celcius_to_farenheit){
    assert_true(celcius_to_farenheit(15) == 59.0, SPOT);
}