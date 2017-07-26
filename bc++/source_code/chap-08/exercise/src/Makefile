
CXX = g++

GTEST_HOME = <GTEST_INSTALL_DIR>
INC = -I $(GTEST_HOME) -I $(GTEST_HOME)/include -I . -I ../src 
LDLIBS = $(GTEST_HOME)/lib/libgtest.a -lstdc++
LDFLAGS = 
CXXFLAGS = -std=c++11 -pthread

TARGET = template_app

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
	$(CXX) -static $(CXXFLAGS) -MMD -c $< -o $@ $(INC)

.PHONY: all clean tests

all: clean $(TARGET) tests

tests: $(TARGET)
	$(BUILD_DIR)/$(TARGET)

clean:
	rm -fr $(BUILD_DIR)

