
#
# g++ -o temp_conversion_test -I gtest-1.7.0/include 
#        -pthrea main.cpp gtest-1.7.0/lib/libgtest.a 
#        -I . -I ../src
#

CXX = g++

GTEST_HOME = <GTEST_INSTALL_DIR> 
INC = -I $(GTEST_HOME) -I $(GTEST_HOME)/include -I . -I ../src
LDLIBS = $(GTEST_HOME)/lib/libgtest.a
LDFLAGS = 
CXXFLAGS = -std=c++11 -pthread

TARGET = temp_conversion_test

BUILD_DIR = ./build
SRCS = $(wildcard *.cpp)
OBJ = $(SRCS:%.cpp=$(BUILD_DIR)/%.o)
DEPS = $(OBJ:%.o=%.d)

$(TARGET) : $(BUILD_DIR)/$(TARGET)

$(BUILD_DIR)/$(TARGET) : $(OBJ)
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

-include $(DEPS)

$(BUILD_DIR)/%.o : %.cpp
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -MMD -c $< -o $@ $(INC)

