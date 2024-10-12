public class TestClass {
    public static int staticVar;
    public int instanceVar;

    public static void main(String[] args) {
        int x = 32768; // ldc
        TestClass myObj = new TestClass(); // new
        TestClass.staticVar = x; // putstatic
        x = TestClass.staticVar; // getstatic
        myObj.instanceVar = x; // putfield
        x = myObj.instanceVar; // getfield
        Object obj = myObj;
        if (obj instanceof TestClass) { // instanceof
            myObj = (TestClass) obj; // checkcast
            System.out.println(myObj.instanceVar);
        }
    }
}