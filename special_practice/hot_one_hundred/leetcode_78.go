package hot_one_hundred

var res [][]int

func Subsets(nums []int) [][]int {
	trace := make([]int, 0)
	BackTrace(0, trace, nums)
	return res
}

func BackTrace(index int, trace []int, nums []int) {
	res = append(res, trace)
	for i := index; i < len(nums); i++ {
		trace = append(trace, nums[i])
		BackTrace(i+1, trace, nums)
		trace = trace[0 : len(trace)-1]
	}
}

/*
class Solution {
public:
    vector<vector<int>> res;
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<int> track;  //记录走过的路径
        backtrack(nums, 0, track);
        return res;
    }
private:
    void backtrack(vector<int>& nums, int start, vector<int>& track){
        //前序遍历
        res.push_back(track);
        //从start开始，防止重复子集产生
        for(int i=start; i<nums.size(); ++i){
            //做选择
            track.push_back(nums[i]);
            //递归回溯
            backtrack(nums, i+1, track);
            //撤销选择
            track.pop_back();
        }
    }
};
*/
