#include <stdio.h>

int* binary_search(int arr[], int size, int item) {

    int low = 0;
    int high = size - 1;

    while (low <= high){
        
        int midd = low + (high - low) / 2;

        if (arr[midd] == item){
            return &arr[midd];
        }

        if (arr[midd] < item){
            low = midd + 1;
        } else {
            high = midd - 1;
        }
    }

    return NULL;
}

int main() {

    int arr[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int size = sizeof(arr) / sizeof(arr[0]);
    int item = 10;

    int* result = binary_search(arr, size, item);
    {
        if (result != NULL) {
            printf("Element %d found at index %d\n", item, result - arr);
        } else {
            printf("Element %d not found\n", item);
        }
    }
}