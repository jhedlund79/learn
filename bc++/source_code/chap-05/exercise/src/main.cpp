
#include <iostream>
#include <string>

#include "functions.hpp"

void generate_frequency_data(const std::string& file_name) {
	initialize();
	calculate(file_name);
	display();
}

int main(int argc, char* argv[]) {

	if (argc != 2) {
		std::cerr << "Missing data file name for calculating frequency of Farmers Markets" << std::endl;
		return 1;
	}

	std::string file_name = argv[1];
	generate_frequency_data(file_name);
}

