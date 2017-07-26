
#include <iostream>
#include <cmath>
#include <string>

/*
	calc <num 1> op <num 2>
		returns the operation on the two numbers
*/

int main(int argc, char* argv[]) {
    if (argc != 4) {
        std::cerr << "Error: invalid number of arguments, must be 4" 
                  << std::endl;
        return 1;
    }

    double lhs = std::stod(argv[1]);
    std::string op = argv[2];
    double rhs = std::stod(argv[3]);

    if (op == "+") {
        std::cout << "result = " << lhs + rhs << std::endl;
    } else if (op == "-") {
        std::cout << "result = " << lhs - rhs << std::endl;
    } else if (op == "x") {
        std::cout << "result = " << lhs * rhs << std::endl;
    } else if (op == "/") {
        std::cout << "result = " << lhs / rhs << std::endl;
    } else if (op == "^") {
        std::cout << "result = " << std::pow(lhs, rhs) << std::endl;
    } else {
        std::cerr << "Error: op " << op << " not supported." << std::endl;
        return 1;
    }
}

