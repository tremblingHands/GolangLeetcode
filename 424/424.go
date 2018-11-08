func characterReplacement(s string, k int) int {
    st, en, maxcnt := 0, 0, 0
    cnt := make([]int, 26)
    for{
        if en == len(s){
            break
        }

        if en-st-maxcnt <= k{
            en++
            cnt[s[en-1]-'A']++
            if cnt[s[en-1]-'A']>maxcnt{
                maxcnt++
            }
        }else{
            cnt[s[st]-'A']--
            st++
        }
    }
    if maxcnt+k > len(s){
        return len(s)
    }
    return maxcnt+k
}
