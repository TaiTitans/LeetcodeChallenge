
class mostwater{
 public int maxArea(int[] height){
    int result = 0;
    int a_pointer = 0;
    int b_pointer = height.length - 1;
    while(a_pointer < b_pointer){
        if(height[a_pointer]<height[b_pointer]){
            result = Math.max(result, height[a_pointer] * (b_pointer-a_pointer));
            a_pointer++;
        } else {
            result = Math.max(result, height[b_pointer] * (b_pointer-a_pointer));
            b_pointer--;
        }
    }
    return result;
 }


public static void main (String[] args){
    mostwater container = new mostwater();
    int[] heights = {1,8,6,2,5,4,8,3,7};
    int maxArea = container.maxArea(heights);
    System.out.println("Max Area: "+ maxArea);;
}
}