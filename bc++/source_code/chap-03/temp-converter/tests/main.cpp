/*
 * main.cpp
 *
 *  Created on: Jul 3, 2015
 *      Author: dubchak
 */

#include <CppUtest/CommandLineTestRunner.h>

#include "fahrenheit_celsius_test.hpp"

int main(int argc, char* argv[]) {
	// our test runner
	return CommandLineTestRunner::RunAllTests(argc, argv);
}
