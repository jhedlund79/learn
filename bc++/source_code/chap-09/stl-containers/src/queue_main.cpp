
#include <iostream>
#include <queue>


int main() {

	std::cout << "Queue" << std::endl;
	std::cout << "=====" << std::endl;

	std::queue<std::string> message_q;

	message_q.push("first");
	message_q.push("second");
	message_q.push("third");

	auto msg = message_q.back();
	std::cout << msg << std::endl;
	message_q.pop();

	message_q.push("fourth");

	std::cout << "Am I empty? " << (message_q.empty() ? "Yes" : "No") << std::endl;
	std::cout << "My size is: " << message_q.size() << std::endl;

}
