/*
 * reverse_test.hpp
 *
 *  Created on: Aug 4, 2015
 *      Author: dubchak
 */

#ifndef TESTS_REVERSE_TEST_HPP_
#define TESTS_REVERSE_TEST_HPP_

#include <algorithm>
#include <functional>
#include <string>

#include <unistd.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <cerrno>

#include "gtest/gtest.h"

std::string reverse(std::string in) {
    for (int i = 0; i < in.length() / 2; ++i) {
        std::string::value_type c = in[i];
        in[i] = in[in.length() - 1 - i];
        in[in.length() - 1 - i] = c;
    }
	return in;
}

TEST(ReverseTest, reverseTestOldSchool) {
	std::string fruit{"apple"};
	EXPECT_STREQ("elppa", reverse(fruit).c_str());
}

TEST(ReverseTest, reverseTestSTL) {
	std::string fruit{"apple"};
	std::reverse(std::begin(fruit), std::end(fruit));
	EXPECT_STREQ("elppa", fruit.c_str());
}


template<typename T>
struct Scale {
private:
   T  _incr;
public:
   explicit Scale(T t) : _incr(t) {}
   T operator()(T arg) {
      return arg * _incr;
   }
};

TEST(ReverseTest, testScale) {

	std::vector<int> a({1, 2, 3, 4, 5});
	std::vector<int> result;
	result.reserve(5);

	Scale<int> s(32);

	std::transform(std::begin(a), std::end(a),
	               std::begin(result), s);

	// result: 32, 64, 96, 128, 160
}

TEST(ReverseTest, testErrno) {

	std::string fname = "/some/path/no_file_found.txt";
	if (open(fname.c_str(), O_RDONLY, S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH ) == -1) {
		std::cerr<< strerror(errno) << std::endl;
	}
}

#endif /* TESTS_REVERSE_TEST_HPP_ */
