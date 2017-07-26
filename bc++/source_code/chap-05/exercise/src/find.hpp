
#ifndef FIND_HPP_
#define FIND_HPP_

#include <string>

std::string __find(std::string value, std::string data[], std::size_t beg, std::size_t end) {

	while (beg <= end) {
        int mid = beg + (end - beg) / 2;
        if (data[mid] == value)
            return value;
        else if (data[mid] < value)
            beg = mid + 1;
        else
            end = mid - 1;
    }

    return ""; // Not found
}

namespace intro {
    std::string (*find)(std::string value, std::string data[], std::size_t beg, std::size_t end) = __find;
}
#endif /* FIND_HPP_ */
