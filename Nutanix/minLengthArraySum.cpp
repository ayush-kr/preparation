//Find minimum length of array with sum >= num

#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int qsum = 0;
        int mini = INT_MAX;
        int left = 0;
        for(int i=0;i<nums.size();i++){
            qsum+=nums[i];
            while(qsum >= s){
                mini = min(mini, i+1-left);
                qsum = qsum-nums[left++];
            }
        }
        if(mini == INT_MAX)
            return 0;
        return mini;
    }
};