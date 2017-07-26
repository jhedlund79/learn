
#ifndef TEST_CELSIUS_TO_FAHRENHEIT_HPP_
#define TEST_CELSIUS_TO_FAHRENHEIT_HPP_

#include "gtest/gtest.h"
#include "temp_conversion.hpp"

TEST(TemperatureConversionTest,  celsius_to_fahrenheit) {
    EXPECT_EQ(59.0, celsius_to_fahrenheit(15));
}

#endif // TEST_CELSIUS_TO_FAHRENHEIT_HPP_
