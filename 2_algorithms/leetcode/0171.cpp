#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
    int titleToNumber(string columnTitle) {
        int num = 0;
        for (char c : columnTitle) {
            num = num * 26 + (c - 'A') + 1;
        }
        return num;
    }
};

// Expose a C-compatible function
extern "C" {
    int titleToNumberWrapper(const char* columnTitle) {
        Solution solution;
        return solution.titleToNumber(string(columnTitle));
    }
}
