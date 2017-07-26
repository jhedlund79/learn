
#include <iostream>
#include <vector>

int main() {

	std::cout << "Vector" << std::endl;
	std::cout << "======" << std::endl;

	std::vector<int> int_months;

	for (int i = 1; i < 13; ++i) {
		int_months.push_back(i);
	}

	// C++
	auto start = int_months.begin();
	while (start != int_months.end()) {
		std::cout << "Value: " << *start << std::endl;
		++start;
	}

	start = std::begin(int_months);
	while (start != std::end(int_months)) {
		std::cout << "Value: " << *start << std::endl;
		++start;
	}

	auto cstart = int_months.cbegin();
	while (cstart != int_months.cend()) {
		std::cout << "Value: " << *cstart << std::endl;
		++cstart;
	}

	std::cout << "Am I empty? " << (int_months.empty() ? "Yes" : "No") << std::endl;
	std::cout << "My size is    : " << int_months.size() << std::endl;
	std::cout << "My capacity is: " << int_months.max_size() << std::endl;

}
