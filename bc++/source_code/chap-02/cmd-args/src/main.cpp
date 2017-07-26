
#include <iostream>

int main(int argc, char* argv[]) {

	if (argc < 6) {
		std::cerr << "Error! Invalid number of arguments. Must be at least 5." << std::endl;
		return 1;
	}

	std::cout << "Total arguments: " << argc << std::endl;
	for (int i = 0; i < argc; ++i) {
		std::cout << "argv[" << i << "] = " << argv[i] << std::endl;
	}
}

