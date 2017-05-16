package main;

import java.util.ArrayList;

/**
 * Created by Bartosz on 16.05.2017
 */
class Medium extends Thread {

    final int SIZE = 500;
    Cable_unit[] cable;
    boolean[] isSending;
    ArrayList<Integer> slots = new ArrayList<>();
    int slot_counter;

    Medium () {
        slot_counter = 0;
        cable = new Cable_unit[SIZE];
        for (int i = 0; i < cable.length; i++) {
            cable[i] = new Cable_unit();
        }
        slots.add(1);
        slots.add(SIZE/2);
        slots.add(SIZE-2);
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
            // przejsc po slotach i sprawdzic czy nadają
            for (int i = 0; i < cable.length; i++) {
                System.out.print(cable[i].value);
            }

            System.out.print("\n");
            try {
                Thread.sleep(2000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            move();
        }
    }

    private void move() {
        Cable_unit[] tmp = new Cable_unit[SIZE];
        for (int i = 0; i < cable.length; i++) {
            
        }
    }
}
