func replaceElements(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        var highest int = -1
        for j := i + 1; j < len(arr); j++ {
            if arr[j] > highest {
                highest = arr[j]
            }
        }
        arr[i] = highest
    }
    return arr
}
