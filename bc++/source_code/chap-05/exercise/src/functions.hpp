
#ifndef FUNCTIONS_HPP_
#define FUNCTIONS_HPP_

#include <string>

#include "file_system.hpp"
#include "string_utils.hpp"

#include "data.hpp"
#include "find.hpp"


constexpr auto aggregate = [&] (int row_num, const std::string& line) {
    auto parts = split(line, "|", true);

    // switch for testing...can you come up with a better way?
    //return intro::find(parts[STATE_POS], states.get(), START_IDX, END_IDX);
    auto val = intro::find(parts[STATE_POS], states.get(), START_IDX, END_IDX);
    if (val != "")
        increment(val);
};

void calculate(const std::string& file_name) {
    fs::FileReader<std::string> reader(file_name);
    if (reader) {
        reader.for_each(aggregate);
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
