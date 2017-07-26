
#ifndef AGGREGATE_TEST_HPP_
#define AGGREGATE_TEST_HPP_

#include <string>
#include "functions.hpp"

TEST(AggregateTest, verifySplitAndFindAggregateFunctionality) {
	initialize();
	std::string line = "1008044|10:10 Farmers Market|http://www.1010farmersmarket.com|5960 Stewart Parkway|Douglasville|Douglas|Georgia|30135|January to December|Thu:4:00 pm - 7:00 pm;|||||||-84.7689|33.7196|Faith-based institution (e.g., church, mosque, synagogue, temple)|Y|N|N|N|N|Y|Y|N|Y|N|N|Y|Y|Y|Y|N|Y|N|N|Y|N|Y|Y|N|N|7/18/2012 13:51";
	std::string actual = aggregate(1, line);
	ASSERT_STREQ("Georgia", actual.c_str());
}

#endif // AGGREGATE_TEST_HPP_