
CXX = g++

GTEST_HOME = <GTEST_INSTALL_DIR>
CLASSLIB_HOME = ../../../class-library
INC = -I $(GTEST_HOME) -I $(GTEST_HOME)/include -I . -I ../src -I $(CLASSLIB_HOME)/include
LDLIBS = $(GTEST_HOME)/lib/libgtest.a -lstdc++
LDFLAGS = 
CXXFLAGS = -std=c++11 -pthread

TARGET = template_app

BUILD_DIR = ./build
SRCS = $(wildcard *.cpp)
OBJ = $(SRCS:%.cpp=$(BUILD_DIR)/%.o)
DEPS = $(OBJ:%.o=%.d)

$(TARGET) : $(BUILD_DIR)/$(TARGET)

stack : $(BUILD_DIR)/stack
	$(BUILD_DIR)/stack
	
array : $(BUILD_DIR)/array
	$(BUILD_DIR)/array
	
vector : $(BUILD_DIR)/vector
	$(BUILD_DIR)/vector
	
map : $(BUILD_DIR)/map
	$(BUILD_DIR)/map
	
queue : $(BUILD_DIR)/queue
	$(BUILD_DIR)/queue
	
deque : $(BUILD_DIR)/deque
	$(BUILD_DIR)/deque

$(BUILD_DIR)/$(TARGET) : $(OBJ)
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

-include $(DEPS)

$(BUILD_DIR)/%.o : %.cpp
	mkdir -p $(@D)
	$(CXX) -static $(CXXFLAGS) -MMD -c $< -o $@ $(INC)

.PHONY: all clean tests

all: $(TARGET)

tests: $(TARGET)
	$(BUILD_DIR)/$(TARGET)

clean:
	rm -fr $(BUILD_DIR)

$(BUILD_DIR)/stack : stack_main.o
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

$(BUILD_DIR)/array : array_main.o
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

$(BUILD_DIR)/vector : vector_main.o
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

$(BUILD_DIR)/map : map_main.o
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

$(BUILD_DIR)/queue : queue_main.o
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)

$(BUILD_DIR)/deque : deque_main.o
	mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS) $(LDLIBS)
