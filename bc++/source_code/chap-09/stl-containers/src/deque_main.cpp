
#include <iostream>
#include <deque>

int main() {

	std::cout << "Deque" << std::endl;
	std::cout << "=====" << std::endl;

	std::deque<int> int_deque;
	int_deque.push_front(10);
	int_deque.push_back(20);
	int_deque.push_front(30);
	int_deque.push_back(40);

	std::cout << "Am I empty? " << (int_deque.empty() ? "Yes" : "No") << std::endl;
	std::cout << int_deque.back() << ", ";
	int_deque.pop_back();

	std::cout << int_deque.back() << ", ";
	int_deque.pop_back();

	std::cout << int_deque.back() << ", ";
	int_deque.pop_back();

	std::cout << int_deque.back() << std::endl;
	int_deque.pop_back();

	std::cout << "Am I empty? " << (int_deque.empty() ? "Yes" : "No") << std::endl;
	std::cout << "My size is    : " << int_deque.size() << std::endl;
	std::cout << "My capacity is: " << int_deque.max_size() << std::endl;

}
