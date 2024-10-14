public class TestClass implements Runnable {
    public static void main(String[] args) {
        new TestClass().test();
    }

    public void test() {
        TestClass.staticMethod(); // invokestatic
        TestClass demo = new TestClass(); // invokespecial
        demo.instanceMethod(); // invokespecial
        super.equals(null); // invokespecial
        this.run(); // invokevirtual
        ((Runnable) demo).run(); // invokeinterface
    }

    public static void staticMethod() {
    }

    private void instanceMethod() {
    }

    @Override
    public void run() {
    }
}