/*
 * functions.hpp
 *
 *  Created on: Jul 21, 2015
 *      Author: dubchak
 */

#ifndef FUNCTIONS_HPP_
#define FUNCTIONS_HPP_

#include <string>

#include "file_system.hpp"
#include "string_utils.hpp"

#include "data.hpp"
#include "find.hpp"

#include "address.hpp"
#include "model.hpp"

data::Site create_site(std::vector<std::string>& data) {
	return data::Site(std::stol(data[0]), data[1],
	    		  addr::WebAddress(data[2]),
				  addr::PostalAddress(data[3], data[4], data[5], to_abbrev(data[6]), data[7]));
}

constexpr auto fx = [&] (int row_num, const std::string& line) {
    auto parts = split(line, "|", true);
    std::string val = intro::find(parts[STATE_POS], states.get(), START_IDX, END_IDX);
    if (val != "") {
    	// Build objects
    	auto s = create_site(parts);
    	std::cout << s.to_string() << std::endl;
    }
};

void calculate(const std::string& file_name) {
    fs::FileReader<std::string> reader(file_name);
    if (reader) {
        reader.for_each(fx);
    }
}

void display() {
    for (auto& state: frequencies) {
        std::cout << std::left
            << std::setw(15)
            << state.first
			<< ": "
            << std::right
            << std::setw(10)
            << state.second
            << std::endl;
    }
}


#endif /* FUNCTIONS_HPP_ */
