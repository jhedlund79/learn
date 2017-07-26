
#include <iostream>
#include <map>

int main() {

	std::cout << "Map" << std::endl;
	std::cout << "===" << std::endl;

	std::map<std::string, std::string> config;

	// using dictionary style
	config["site.name"] = "www.google.com";

	// using std::make_pair function
	config.insert(std::make_pair("site.name", "www.google.com"));

	std::cout << config["site.name"] << std::endl;

	auto iter = config.find("missing.key");
	if (iter != std::end(config))
		std::cout << "Key 'missing.key' was found" << std::endl;
	else
		std::cout << "Key 'missing.key' was not found" << std::endl;

	std::cout << "Am I empty? " << (config.empty() ? "Yes" : "No") << std::endl;
	std::cout << "My size is    : " << config.size() << std::endl;
	std::cout << "My capacity is: " << config.max_size() << std::endl;

}
