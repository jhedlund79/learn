/*
 * site.hpp
 *
 *  Created on: Jul 26, 2015
 *      Author: dubchak
 */

#ifndef SITE_HPP_
#define SITE_HPP_

#include <string>

#include "address.hpp"
#include "_us_address.hpp"
#include "_web_address.hpp"


class _site {
public:
	_site() : _id(-1) {}
	_site(const _site& rhs) : _id(rhs._id), _name(rhs._name), _uri(rhs._uri), _postal(rhs._postal) {}
	_site(long id, std::string name,
			_web_address uri, _us_address postal)
				: _id(id), _name(name), _uri(uri), _postal(postal) {}
	virtual ~_site() {}

	_site& operator=(const _site& rhs)  {
		if (&rhs != this) {
			_id = rhs._id;
			_name = rhs._name;
			_uri = rhs._uri;
			_postal = rhs._postal;
		}
		return *this;
	}

	std::string to_string() const { return _name + ": " + _uri.to_string(); }
	std::string name() const { return _name; }
	std::string state() const { return ::to_string(_postal.state()); }
	std::string city() const { return _postal.city(); }


private:
	long			_id;
	std::string		_name;
	_web_address	_uri;
	_us_address		_postal;

};

template<>
bool operator<(_site& lhs, _site& rhs) {
	return lhs.name() < rhs.name();
}

#endif /* SITE_HPP_ */
