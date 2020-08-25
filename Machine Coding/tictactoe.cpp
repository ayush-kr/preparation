#include<bits/stdc++.h>
using namespace std;
#define MAX 10

void draw(int player, int r, int c);
bool checkEndOfGame(int r, int c);
void announceWinner(char c);
void display();
void init(int len);
void startGame();
bool isValidInput(int num);
bool isValidMove(int r, int c);
bool isValid(int num);

int n;
char matrix[MAX][MAX];
int empty;

char letter[2] = {'0', 'X'};

bool isValidInput(int num){
    if(num<=0 || num > n*n) return false;
    return true;
}

bool isValidMove(int r, int c){
    if (matrix[r][c] != '-'){
        cout<<"Invalid Input. Please Try Again"<<endl;
        return false;
    } 
    return true;
}

bool isValid(int num){
    if (!isValidInput(num)){
        cout<<"Invalid Input. Please Try Again"<<endl;
        return false;
    };
    int r = (num-1)/n;
    int c = (num-1)%n;
    return isValidMove(r, c);
}

void draw(int player, int r, int c){
    matrix[r][c] = letter[player];
    //Decrement number of empty spaces by 1
    empty--;
}

bool checkEndOfGame(int r, int c){
    char ch = matrix[r][c];
    //check horizontal
    for(int i=0;i<n;i++){
        if(ch != matrix[r][i]) break;
        if(i == n-1){
            announceWinner(ch);
            return true;
        }
    }

    //check Verical
    for(int i=0;i<n;i++){
        if(ch != matrix[i][c]) break;
        if(i == n-1){
            announceWinner(ch);
            return true;
        }
    }

    //check 1st Diagonal that goes from top right to bottom left
    if (r + c == n - 1){
        for(int i=0;i<n;i++){
            if(matrix[i][n-i-1] != ch) break;
            if(i == n-1){
                announceWinner(ch);
                return true;
            }
        }
    }

    //check 2nd Diagonal that goes from top left to bottom right
    if(r == c){
        for(int i=0;i<n;i++){
            if(matrix[i][i] != ch)break;
            if(i == n-1){
                announceWinner(ch);
                return true;
            }
        }
    }

    //Check Draw case
    if (empty == 0) {
        cout<<"Game Over: Draw"<<endl;
        return true;
    }
    return false;
}

void announceWinner(char c){
    cout<<"Game Over: Player "<<c<<" has won"<<endl;
}

void display(){
    for(int i=0;i<n;i++){
        for(int j=0;j<n;j++)
            cout<<matrix[i][j]<<" ";
        cout<<endl;
    }
}
void init(int len){
    n = len;
    //Initialise empty spaces by total number of cells
    empty = n*n;

    for(int i=0;i<n;i++)
        for(int j=0;j<n;j++)
            matrix[i][j] = '-';
}

void startGame(){

    init(3);

    int player = 1;
    int r,c;
    do{
        display();
        player = !player;

        //taking input number as string
        string number;      
        do{
            cout<<"Player "<<letter[player]<<"'s turn, Enter any number between 1 to "<<n*n<<endl;
            cin>>number;
        } while(!isValid(atoi(number.c_str()))); //convert the number to integer and check validity

        int index = atoi(number.c_str());
        
        r = (index-1)/n;
        c = (index-1)%n;

        draw(player, r,c);
    }while(!checkEndOfGame(r, c));
}

int main(){
    startGame();
    return 0;
}