package main;

import java.util.ArrayList;

/**
 * Created by Bartosz on 16.05.2017
 */
class Medium extends Thread {

    private final int SIZE = 500;
    String[] cable;
    boolean[] isSending;
    ArrayList<Integer> slots = new ArrayList<>();
    int slot_counter;

    Medium () {
        slot_counter = 0;
        cable = new String[SIZE];
        for (int i = 0; i < cable.length; i++) {
            cable[i] = "-";
        }
        slots.add(0);
        slots.add(SIZE/2);
        slots.add(SIZE-1);
        isSending = new boolean[slots.size()];
        for( boolean b : isSending) b = false;
    }

    int getSlot() {
        int ret = -1;
        if(slot_counter < slots.size()) {
            ret = slots.get(slot_counter);
            slot_counter++;
        }
        return ret;
    }

    public void run() {
        while(true) {
            int counter = 0;
            for (int i = 0; i < cable.length; i++) {
                for (int j : slots) {
                    if (i == j) {
                        if (isSending[counter]) {
                            cable[i] = String.valueOf(counter);
                        }

                        counter++;
                    }
                }
                System.out.print(cable[i]);
            }
            System.out.print("\n");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
