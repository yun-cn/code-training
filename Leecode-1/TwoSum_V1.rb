def two_sum(nums, target) 
    seach = Hash.new
    nums.each_with_index do |num, index|
        i = search[target - num]
        return [index, i] if i != nil
        search[num] = index
    end
end