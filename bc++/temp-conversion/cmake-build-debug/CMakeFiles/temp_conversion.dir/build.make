# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.7

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /Applications/CLion.app/Contents/bin/cmake/bin/cmake

# The command to remove a file.
RM = /Applications/CLion.app/Contents/bin/cmake/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/jasonhedlund/workspace/courses/bc++/temp-conversion

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/temp_conversion.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/temp_conversion.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/temp_conversion.dir/flags.make

CMakeFiles/temp_conversion.dir/src/main.cpp.o: CMakeFiles/temp_conversion.dir/flags.make
CMakeFiles/temp_conversion.dir/src/main.cpp.o: ../src/main.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/temp_conversion.dir/src/main.cpp.o"
	/Library/Developer/CommandLineTools/usr/bin/c++   $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/temp_conversion.dir/src/main.cpp.o -c /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/src/main.cpp

CMakeFiles/temp_conversion.dir/src/main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/temp_conversion.dir/src/main.cpp.i"
	/Library/Developer/CommandLineTools/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/src/main.cpp > CMakeFiles/temp_conversion.dir/src/main.cpp.i

CMakeFiles/temp_conversion.dir/src/main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/temp_conversion.dir/src/main.cpp.s"
	/Library/Developer/CommandLineTools/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/src/main.cpp -o CMakeFiles/temp_conversion.dir/src/main.cpp.s

CMakeFiles/temp_conversion.dir/src/main.cpp.o.requires:

.PHONY : CMakeFiles/temp_conversion.dir/src/main.cpp.o.requires

CMakeFiles/temp_conversion.dir/src/main.cpp.o.provides: CMakeFiles/temp_conversion.dir/src/main.cpp.o.requires
	$(MAKE) -f CMakeFiles/temp_conversion.dir/build.make CMakeFiles/temp_conversion.dir/src/main.cpp.o.provides.build
.PHONY : CMakeFiles/temp_conversion.dir/src/main.cpp.o.provides

CMakeFiles/temp_conversion.dir/src/main.cpp.o.provides.build: CMakeFiles/temp_conversion.dir/src/main.cpp.o


CMakeFiles/temp_conversion.dir/test/test.cpp.o: CMakeFiles/temp_conversion.dir/flags.make
CMakeFiles/temp_conversion.dir/test/test.cpp.o: ../test/test.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/temp_conversion.dir/test/test.cpp.o"
	/Library/Developer/CommandLineTools/usr/bin/c++   $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/temp_conversion.dir/test/test.cpp.o -c /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/test/test.cpp

CMakeFiles/temp_conversion.dir/test/test.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/temp_conversion.dir/test/test.cpp.i"
	/Library/Developer/CommandLineTools/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/test/test.cpp > CMakeFiles/temp_conversion.dir/test/test.cpp.i

CMakeFiles/temp_conversion.dir/test/test.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/temp_conversion.dir/test/test.cpp.s"
	/Library/Developer/CommandLineTools/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/test/test.cpp -o CMakeFiles/temp_conversion.dir/test/test.cpp.s

CMakeFiles/temp_conversion.dir/test/test.cpp.o.requires:

.PHONY : CMakeFiles/temp_conversion.dir/test/test.cpp.o.requires

CMakeFiles/temp_conversion.dir/test/test.cpp.o.provides: CMakeFiles/temp_conversion.dir/test/test.cpp.o.requires
	$(MAKE) -f CMakeFiles/temp_conversion.dir/build.make CMakeFiles/temp_conversion.dir/test/test.cpp.o.provides.build
.PHONY : CMakeFiles/temp_conversion.dir/test/test.cpp.o.provides

CMakeFiles/temp_conversion.dir/test/test.cpp.o.provides.build: CMakeFiles/temp_conversion.dir/test/test.cpp.o


# Object files for target temp_conversion
temp_conversion_OBJECTS = \
"CMakeFiles/temp_conversion.dir/src/main.cpp.o" \
"CMakeFiles/temp_conversion.dir/test/test.cpp.o"

# External object files for target temp_conversion
temp_conversion_EXTERNAL_OBJECTS =

temp_conversion: CMakeFiles/temp_conversion.dir/src/main.cpp.o
temp_conversion: CMakeFiles/temp_conversion.dir/test/test.cpp.o
temp_conversion: CMakeFiles/temp_conversion.dir/build.make
temp_conversion: CMakeFiles/temp_conversion.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Linking CXX executable temp_conversion"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/temp_conversion.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/temp_conversion.dir/build: temp_conversion

.PHONY : CMakeFiles/temp_conversion.dir/build

CMakeFiles/temp_conversion.dir/requires: CMakeFiles/temp_conversion.dir/src/main.cpp.o.requires
CMakeFiles/temp_conversion.dir/requires: CMakeFiles/temp_conversion.dir/test/test.cpp.o.requires

.PHONY : CMakeFiles/temp_conversion.dir/requires

CMakeFiles/temp_conversion.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/temp_conversion.dir/cmake_clean.cmake
.PHONY : CMakeFiles/temp_conversion.dir/clean

CMakeFiles/temp_conversion.dir/depend:
	cd /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/jasonhedlund/workspace/courses/bc++/temp-conversion /Users/jasonhedlund/workspace/courses/bc++/temp-conversion /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug /Users/jasonhedlund/workspace/courses/bc++/temp-conversion/cmake-build-debug/CMakeFiles/temp_conversion.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/temp_conversion.dir/depend

