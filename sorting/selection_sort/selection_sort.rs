
fn main(){
    selection_sort([1,4,3,7,2,6])
}

fn selection_sort<T>(mut arr: &mut [T]) where T: Ord {
    for i in 0..arr.len() - 1 {
        let mut min_index = i;
        for j in i + 1..arr.len() {
            if arr[j] < arr[min_index] {
                min_index = j;
            }
        }
        arr.swap(i, min_index);
    }
}

fn selection_sort_type_2<T>(mut arr: &mut [T]) where T: Ord {
    arr.iter_mut()
        .enumerate()
        .min_by(|&(i, a), &(j, b)| a.cmp(b))
        .map(|(i, _)| i)
        .and_then(|min_index| arr.get_mut(min_index))
        .and_then(|min_element| arr.get_mut(0).swap(min_element));
}
