#include <iostream>

using namespace std;

int binary_search(int arr[], int size, int item) {
    int low = 0;
    int high = size - 1;

    while (low <= high) {
        int midd = low + (high - low) / 2;
        int guess = arr[midd];

        if (guess == item) {
            return midd; 
        }

        if (guess < item) {
            low = midd + 1;
        } else {
            high = midd - 1; 
        }
    }

    return -1;
}

int main() {
    int arr[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int item = 4;

    int res = binary_search(arr, 10, item);

    cout << res << endl; 

    return 0;
}