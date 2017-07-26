/*
 * functions.hpp
 *
 *  Created on: Jul 21, 2015
 *      Author: dubchak
 */

#ifndef FUNCTIONS_HPP_
#define FUNCTIONS_HPP_

#include <algorithm>
#include <map>
#include <string>
#include <vector>

#include "file_system.hpp"
#include "string_utils.hpp"

#include "data.hpp"
#include "find.hpp"

#include "address.hpp"
#include "model.hpp"

// Key aliases
using Key = std::string;
using CityKey = Key;
using StateKey = Key;

// Market aliases
using Market = data::Site;
using Markets = std::vector<Market>;
using CityMarkets = std::map<CityKey, Markets>;
using StateMarkets = std::map<StateKey, CityMarkets>;


Market create_site(std::vector<std::string>& data) {

	return Market(std::stol(data[0]), data[1],
			addr::WebAddress(data[2]),
			addr::PostalAddress(data[3], data[4], data[5], to_abbrev(data[6]), data[7]));
}

CityMarkets create_city_market(const Market& market) {
	Markets m;
	m.push_back(market);
	CityMarkets cm;
	cm.insert(std::make_pair(market.city(), m));
	return cm;
}

struct MarketStats {

	MarketStats() {}
	MarketStats(const MarketStats& rhs) : _sites(rhs._sites) {}
	virtual ~MarketStats() {}

	MarketStats& operator=(const MarketStats& rhs) {
		if (&rhs != this) {
			_sites = rhs._sites;
		}
		return *this;
	}

	void add(const Market& s) { _sites.emplace_back(s); }
	Markets& sites() { return _sites; }

	typename Markets::size_type market_count() const { return _sites.size(); }

	// State member functions
	bool contains_state(const StateKey& key) const { return std::find(std::begin(_states), std::end(_states), key) != std::end(_states); }
//	bool contains_city(const Market& market) const { return std::find(std::begin(_cities), std::end(_cities), [&] (std::pair<CityKey, CityMarkets>& p) {
//		return market.city() != p.first; }) != std::end(_cities); }
	void add_market(const Market& market) {
    	if (!contains_state(market.state())) {
    		_states.insert(std::make_pair(market.state(), create_city_market(market)));
		} else {
//			if (!contains_city(market)) {
//				CityMarkets cm;
//				auto cities = _states[market.state()][market.city()];
//				cities.push_back(market);
//			} else if (!contains_state(market.state())) {
//
//			}
		}

	}

	//bool operator!=(const StateMarkets& markets, const CityMarkets& city)

private:
	Markets _sites;
	StateMarkets _states;
	CityMarkets _cities;
};

bool operator==(const std::pair<StateKey, CityMarkets>& p, const std::string& city) {
	return p.second.find(city) == std::end(p.second);
}

std::vector<int> b({1, 2, 3, 4, 5});

struct MarketParser {

	void operator()(int row_num, const std::string& line) {
		auto parts = split(line, "|", true);
		std::string val = intro::find(parts[STATE_POS], states.get(), START_IDX, END_IDX);
		if (to_abbrev(val) != StateAbbrev::INVALID)
			_stat.add(create_site(parts));
	}

	MarketStats& data() { return _stat; }


private:
	MarketStats _stat;
};

struct SiteStateComparator {
	bool operator()(data::Site& lhs, data::Site& rhs) {
		return lhs.state() < rhs.state();
	}
};


void calculate(std::string file_name) {

	MarketParser parser;
    fs::FileReader<std::string> reader(file_name);
    if (reader) {
        reader.for_each(parser);
    }

    MarketStats stats(parser.data());
    std::cout << "There were " << stats.market_count() << " markets read." << std::endl;

    SiteStateComparator comp;
    std::sort(std::begin(stats.sites()), std::end(stats.sites()), comp);
    std::string previous;
    for (auto& m: stats.sites()) {
    	if (previous == "")
    		previous = m.state();

    	stats.add_market(m);
    }
}

void display() {
    for (auto& state: frequencies) {
        std::cout << std::left
            << std::setw(15)
            << state.first
			<< ": "
            << std::right
            << std::setw(10)
            << state.second
            << std::endl;
    }
}


#endif /* FUNCTIONS_HPP_ */
