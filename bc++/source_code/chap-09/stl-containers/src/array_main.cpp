
#include <iostream>
#include <array>


int main() {

	std::cout << "Array" << std::endl;
	std::cout << "======" << std::endl;

	std::array<int, 12> int_months;

	for (int i = 1; i < 13; ++i) {
		int_months[i - 1] = i;
	}

	for (int i = 0; i < 12; ++i) {
		int_months[i] = i + 1;
	}

	for (int i = 0; i < int_months.size(); ++i) {
		std::cout << int_months[i] << std::endl;
	}

	std::cout << "Am I empty? " << (int_months.empty() ? "Yes" : "No") << std::endl;
	std::cout << "My size is    : " << int_months.size() << std::endl;
	std::cout << "My capacity is: " << int_months.max_size() << std::endl;

}
