# Longest Common Prefix

Link: https://leetcode.com/problems/longest-common-prefix/description/

The original constraint says that a string contains only English alphabet.

I've decided to make the task more difficult and my strings can contain any symbol.

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

## Example 1:

Input: strs = ["flower","flow","flight"]

Output: "fl"

## Example 2:

Input: strs = ["dog","racecar","car"]

Output: ""

Explanation: There is no common prefix among the input strings.

## Constraints:

1 <= strs.length <= 200

0 <= strs[i].length <= 200

strs[i] consists of only lowercase English letters if it is non-empty.
