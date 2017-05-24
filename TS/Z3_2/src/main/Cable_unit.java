package main;

/**
 * Created by Bartosz on 16.05.2017
 */
class Cable_unit {

    private Host master;
    private int position;
    private int signalDuration;
    private Direction direction;
    private boolean basic = true;

    Cable_unit(Host h, int position, int signalDuration, Direction dir) {
        this.master = h;
        this.position = position;
        this.signalDuration = signalDuration;
        this.direction = dir;
    }

    void tick() {
        signalDuration--;

        if(basic){
            basic = false;
            return;
        }

        if(direction==Direction.RIGHT){
            position++;

            if(position >= Ethernet.getLength()){
                signalDuration = 0;
                position--;
            }
        } else {
            position--;
            if(position<0){
                signalDuration = 0;
                position++;
            }
        }
    }

    public int getSignalDuration() {
        return signalDuration;
    }

    public Host getMaster() {
        return master;
    }

    public int getPosition() {
        return position;
    }
    
    Direction getDirection(){
        return direction;
    }

    public void lost() {
        master = null;
    }
}
