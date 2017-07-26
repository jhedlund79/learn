/*
 * fahrenheit_celcius_test.hpp
 *
 *  Created on: Jul 3, 2015
 *      Author: dubchak
 */

#ifndef TESTS_FAHRENHEIT_CELCIUS_TEST_HPP_
#define TESTS_FAHRENHEIT_CELCIUS_TEST_HPP_

#include <CppUtest/TestHarness.h>
#include "fahrenheit_to_celsius.hpp"

TEST_GROUP(TempConverterGroup) {

};

TEST(TempConverterGroup, ConvertCelciusToFarhenheit) {
	double fahrenheit = 50;
	double expected = 10;
	CHECK(expected == fahrenheit_to_celsius(fahrenheit));
}


#endif /* TESTS_FAHRENHEIT_CELCIUS_TEST_HPP_ */
