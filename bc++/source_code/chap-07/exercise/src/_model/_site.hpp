/*
 * site.hpp
 *
 *  Created on: Jul 26, 2015
 *      Author: dubchak
 */

#ifndef SITE_HPP_
#define SITE_HPP_

#include <string>

#include "_us_address.hpp"
#include "_web_address.hpp"


class _site {
public:
	_site() : _id(-1) {}
	_site(long id, std::string name,
			_web_address uri, _us_address postal)
				: _id(-1), _name(name), _uri(uri), _postal(postal) {}
	virtual ~_site() {}

	std::string to_string() const { return _name + ": " + _uri.to_string(); }

private:
	long			_id;
	std::string		_name;
	_web_address	_uri;
	_us_address		_postal;

};


#endif /* SITE_HPP_ */
