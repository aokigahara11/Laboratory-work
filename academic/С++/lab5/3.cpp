#include <iostream>
#include <cstdlib>
#include <ctime>
#include <algorithm>

using namespace std;

// Дан двумерный массив. Поменять местами значения элементов столбца и строки на месте стыка
// минимального значения массива (или первого из минимальных).

int main() {
    int rows, cols;
    
    cout << "Введите количество строк и столбцов (должны быть равны друг другу): ";
    cin >> rows >> cols;

    srand(time(NULL));
    int arr[rows][cols]; 

    cout << "\nИсходный массив:\n";
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            arr[i][j] = rand() % 100;
            cout << arr[i][j] << "\t";
        }
        cout << "\n";
    }

    int min_i = 0, min_j = 0;
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            if (arr[i][j] < arr[min_i][min_j]) {
                min_i = i;
                min_j = j;
            }
        }
    }
    
    cout << "\nМинимальный элемент: arr[" << min_i << "][" << min_j << "] = " << arr[min_i][min_j] << "\n";

    for (int k = 0; k < rows; k++) {
        int temp = arr[min_i][k];
        arr[min_i][k] = arr[k][min_j];
        arr[k][min_j] = temp;
    }

    cout << "\nМассив после замены:\n";
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            cout << arr[i][j] << "\t";
        }
        cout << "\n";
    }

    return 0;
}