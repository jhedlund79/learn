/*
 * intro.hpp
 *
 *  Created on: Jul 26, 2015
 *      Author: dubchak
 */

#ifndef MODEL_HPP_
#define MODEL_HPP_

#include "_model/_algorithm.hpp"
#include "_model/_email_address.hpp"
#include "_model/_site.hpp"
#include "_model/_us_address.hpp"
#include "_model/_web_address.hpp"

namespace addr {

	using EmailAddress = _email_address;
	using WebAddress = _web_address;
	using PostalAddress = _us_address;

}

namespace data {
	using Site = _site;
}

#endif /* MODEL_HPP_ */
