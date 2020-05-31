//Given a vector, sort in such a way that -ves remain together and +ves remain together in order they come

#include<bits/stdc++.h>
using namespace std;
class Solution{
    public:
        void sortOrder(vector<int> arr){
            for(int i=0;i<arr.size();i++){
                if(arr[i]<0){
                    for(int j=i;j>0 && arr[j-1]>0;j--){
                        int val = arr[j-1];
                        arr[j-1] = arr[j];
                        arr[j] = val;
                    }
                }
            }
            for(int i=0;i<arr.size();i++){
                cout<<arr[i]<<" ";
            }
            cout<<endl;
        }
        
};
int main(){
    vector<int>input = {1, -2, 3, 4, -5, 2, -8};
    Solution s;
    s.sortOrder(input);
    return 0;
}
