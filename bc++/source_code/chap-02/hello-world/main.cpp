
#include <iostream>
#include <string>

constexpr auto greeting = "hello world!";
auto fx = [] (auto msg) { return std::string("hello ") + msg; };

int main() {

    std::cout << greeting << std::endl;
    std::cout << fx("John") << std::endl;

}

