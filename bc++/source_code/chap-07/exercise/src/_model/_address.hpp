/*
 * _adress.hpp
 *
 *  Created on: Jul 26, 2015
 *      Author: dubchak
 */

#ifndef _ADDRESS_HPP_
#define _ADDRESS_HPP_

#include <string>

class _address {
public:
	_address() {}
	_address(const _address& rhs) : _street1(rhs._street1), _street2(rhs._street2), _city(rhs._city) {}
	_address(std::string s1, std::string s2, std::string c)
			: _street1(s1), _street2(s2), _city(c) {}
	virtual ~_address() {}

	_address& operator=(const _address& rhs) {
		if (&rhs != this) {
			_street1 = rhs._street1;
			_street2 = rhs._street2;
			_city = rhs._city;
		}

		return *this;
	}

	virtual std::string street1() const { return _street1; }
	virtual std::string street2() const { return _street2; }
	virtual std::string city() const { return _city; }

	virtual std::string to_string() const = 0;

private:
	std::string _street1;
	std::string _street2;
	std::string _city;
};


#endif /* _ADDRESS_HPP_ */
