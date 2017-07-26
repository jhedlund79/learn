#include <iostream>
#include <cmath>

int main(int argc, char* argv[]) {
    if (argc != 4) {
        std::cerr << "Error. Incorrect number of arguments"
                  << std::endl;
        return 1;
    }
    double lhs = std::stod(argv[1]);
    std::string op = (argv[2]);
    double rhs = std::stod(argv[3]);

    if (op == "+"){
        std::cout << "result = " << lhs + rhs << std::endl;
    }else if (op == "-"){
        std::cout << "result = " << lhs - rhs << std::endl;
    }else if (op == "/"){
        std::cout << "result = " << lhs / rhs << std::endl;
    }else if (op == "*"){
        std::cout << "result = " << lhs  *rhs << std::endl;
    }else if (op == "^"){
        std::cout << "result = " << std::pow(lhs, rhs) << std::endl;
    }else{
        std::cerr << "Error. op " << op << "not supported."
                  << std::endl;
            return 1;
    }
}
