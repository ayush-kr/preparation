#include<bits/stdc++.h>
using namespace std;

class Board {
    vector<vector<int> > mat = vector<vector<int> >(4, vector<int>(4,0));
    public:

    Board(){
        setRandom();
        display();
    }

    void setRandom(){
        int i = rand()%4;
        int j = rand()%4;
        while(mat[i][j]){
            i = rand()%4;
            j = rand()%4;
        }
        mat[i][j] = rand()%2 ? 2 : 4;
    }

    vector<int> getRowCol(bool isRow, int j, vector<vector<int>> &mat){
        vector<int>arr;
        for(int i = 0; i < 4; i++)
            arr.push_back(isRow ? mat[j][i] : mat[i][j]);
        return arr;
    }

    void setRowCol(bool isRow, int j, vector<int>&arr, vector<vector<int>> &mat){
        for(int i = 0; i < 4; i++){
            if (isRow) mat[j][i] = arr[i];
            else mat[i][j] = arr[i];
        }
    }

    void compress(vector<int>&arr){
        for(int i = 0, j = 0; i < 4; i++){
            if(arr[j] == 0){
                arr.erase(arr.begin() + j);
                arr.push_back(0);
            }else j++;
        }
        for(int i = 0;i < 3; i++){
            if(arr[i] == arr[i+1]){
                arr[i] *= 2;
                arr.erase(arr.begin() + i + 1);
                arr.push_back(0);
            } 
        }
    }

    void reverseCompress(vector<int>&arr){
        reverse(arr.begin(), arr.end());
        compress(arr);
        reverse(arr.begin(), arr.end());
    }

    vector<vector<int> > move(char direction, vector<vector<int> > matrix){
        switch (direction)
        {
        case 'w':
            for(int i = 0; i < 4; i++){
                vector<int> arr = getRowCol(false, i, matrix);
                compress(arr);
                setRowCol(false, i, arr, matrix);
            }
            break;
        case 's':
            for(int i = 0; i < 4; i++){
                vector<int> arr = getRowCol(false, i, matrix);
                reverseCompress(arr);
                setRowCol(false, i, arr, matrix);
            }
            break;
        case 'a':
            for(int i = 0; i < 4; i++){
                vector<int> arr = getRowCol(true, i, matrix);
                compress(arr);
                setRowCol(true, i, arr, matrix);
            }
            break;
        case 'd':
            for(int i = 0; i < 4; i++){
                vector<int> arr = getRowCol(true, i, matrix);
                reverseCompress(arr);
                setRowCol(true, i, arr, matrix);
            }
            break;
        default:
            cout<<"Invalid move! Enter w, a, s, d";
            break;
        }
        return matrix;
    }

    bool makeMove(char direction){
        if(!isStateChanged(move(direction, mat))){
            display();
            return true;
        }
        mat = move(direction, mat);
        setRandom();
        display();
        if(hasLost()){
            cout<<"You Lose!"<<endl;
            return false;
        }
        return true;
    }

    bool hasLost(){
        string moves = "wasd";
        for(auto c : moves)
            if(isStateChanged(move(c, mat)))
                return false;
        return true;
    }

    bool isStateChanged(vector<vector<int> > matrix){
        for(int i = 0; i < 4; i++)
            for(int j = 0; j < 4; j++)
                if(mat[i][j] != matrix[i][j])
                    return true;
        return false;
    }

    void display(){
        for(auto x : mat){
            for(auto y : x) cout<<y<<" ";
            cout<<endl;
        }
        cout<<endl;
    }
};


int main(){
    cout<<"Game starts!"<<endl;
    Board board = Board();
    char direction;
    cin>>direction;

    bool canProceed = board.makeMove(direction);
    while(canProceed){
        cin>>direction;
        canProceed = board.makeMove(direction);
    }
    cout<<"Game Over!"<<endl;
    return 0;
}

