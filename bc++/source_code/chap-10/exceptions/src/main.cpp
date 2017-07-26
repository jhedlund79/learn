
#include <iostream>
#include <string>

#include "file_system.hpp"

constexpr auto fx = [&] (int rownum, const std::string& line) {

};

int main(int argc, char* argv[]) {

	try {
		fs::FileReader<std::string> reader("some_fake_file.txt");
		reader.for_each(fx);
	} catch (const fs::FileReaderException& e) {
		std::cerr << std::endl << "ERROR: " << e.what() << std::endl;
	}

}
