/*
 * email_address.hpp
 *
 *  Created on: Jul 26, 2015
 *      Author: dubchak
 */

#ifndef EMAIL_ADDRESS_HPP_
#define EMAIL_ADDRESS_HPP_

#include <string>

#include "_address.hpp"

class _email_address : public _address {
public:
	_email_address() {}
	_email_address(const _email_address& rhs) : _address(rhs), _email(rhs._email) {}
	_email_address(std::string e) : _email(e) {}
	virtual ~_email_address() {}

	_email_address& operator=(const _email_address& rhs) {
		_address::operator=(rhs);

		if (&rhs != this) {
			_email = rhs._email;
		}

		return *this;
	}

	std::string to_string() const override {
		return _email;
	}

private:
	std::string _email;
};


#endif /* EMAIL_ADDRESS_HPP_ */
