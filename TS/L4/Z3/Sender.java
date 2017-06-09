import java.net.*;
import java.util.*;
import java.util.concurrent.CountDownLatch;

class Sender {


	private static final int datagramSize = 50;
	private static final int sleepTime = 500;
	private static final int maxPacket = 50;
	//ILOSC PAKIETOW KTORE MOZNA WYSLAC BEZ SPRAWDZANIA 
	private static final int retsransmissionWindow = 7;

	//CZAS CZEKANIA NA RETRANSMISJE
	private static final int reSleep = sleepTime * retsransmissionWindow + 5000;
	private InetAddress localHost;
	private int destinationPort;
	private DatagramSocket socket;

	private SenderThread sender;
	private ReceiverThread receiver;
	private ResendThread resender;
	
	// Elementy wyslane
	private ArrayList<Byte> orgData = new ArrayList<Byte>();

	// ZMIENNA POTRZEBNA DO SYNCHRONIZACJI SENDERTHREAD I RESENDTHREAD
	private CountDownLatch latch = new CountDownLatch(1);

	private int packetSend = 0;
	private int packetReceived = 0;
	private DatagramPacket lastPacket;

	//ZMIENNA POTRZEBNA DO KONTROLOWANIA WATKA SENDERTHREAD
	private boolean goNext = true;


	private Sender(int myPort, int destPort) throws Exception {
		
		localHost = InetAddress.getByName("127.0.0.1");
	    destinationPort = destPort;
	    socket = new DatagramSocket(myPort);

	    sender = new SenderThread();
	    receiver = new ReceiverThread();
	    resender = new ResendThread();
	}

	public static void main(String[] args) throws Exception {
		Sender sender = new Sender( Integer.parseInt(args[0]), Integer.parseInt(args[1]));
		sender.sender.start();
	    sender.receiver.start();
	    sender.resender.start();
	}

	class SenderThread extends Thread {
		public void run() {
		    int i, x;
		    int packetInRow = 1;
		    try {

			    for(i=0; (x=System.in.read()) >= 0 ; i++) {
		            
		            Z2Packet p = new Z2Packet(4+1);
		            p.setIntAt(i,0);
		            p.data[4] = (byte) x;
				    DatagramPacket packet = 
					new DatagramPacket(p.data, p.data.length, 
							   localHost, destinationPort);
				    socket.send(packet);
				    System.out.println("S: "+p.getIntAt(0)+": "+ (char) p.data[4]);
				    orgData.add((byte)x);
				    packetSend++;
				    if(packetInRow == retsransmissionWindow) {

				    	latch.await();
				    	latch = new CountDownLatch(1);
				    	packetInRow = 0;
				    }
				    packetInRow++;
				    sleep(sleepTime);
				    
				    while(!goNext){
				    	sleep(sleepTime);
				    }
				}
			
			}
		    catch(Exception e)
			{
		        System.out.println("Z2Sender.SenderThread.run: ");
		        e.printStackTrace();
			}
		}
	}

	class ReceiverThread extends Thread {

		public void run() {
		    try
			{
			    while(true)
				{
				    byte[] data = new byte[datagramSize];
				    DatagramPacket packet =	new DatagramPacket(data, datagramSize);
				    socket.receive(packet);
				    packetReceived++;

		            Z2Packet p = new Z2Packet(packet.getData());
		            if(lastPacket == null){
		            	lastPacket = packet;
		            } else {
		            	Z2Packet l = new Z2Packet(lastPacket.getData());
		            	if(p.getIntAt(0)>l.getIntAt(0))
		            	lastPacket = packet;
		        	}
				    System.out.println("r: "+p.getIntAt(0)+
						       ": "+(char) p.data[4]);
				}
			}
		    catch(Exception e)
			{
		        System.out.println("Z2Sender.ReceiverThread.run: ");
		        e.printStackTrace();
			}
		}

	}

	class ResendThread extends Thread {
		public void run(){
			try{
				while(true){
					if(goNext)
						sleep(sleepTime * retsransmissionWindow);
					sleep(reSleep);
					System.out.println("Sprawdzam");
					Z2Packet pakiet;
					if(lastPacket != null){
						pakiet = new Z2Packet(lastPacket.getData());
					} else {
						pakiet = new Z2Packet(4 + 1);
						pakiet.setIntAt(0,0);
						pakiet.data[4] = (byte) 'X';
					}
					if( (pakiet.getIntAt(0)+1) >= packetSend) {
						latch.countDown();
						goNext = true;
						continue;
					} else {
						goNext = false;
						if(lastPacket != null){
							Z2Packet p = new Z2Packet(lastPacket.getData());
							int id = p.getIntAt(0);
							for (int i = id + 1; i < orgData.size(); i++){
								Z2Packet pa = new Z2Packet(4+1);
		       			    	pa.setIntAt(i,0);
		            			pa.data[4] = (byte) orgData.get(i);
				    			DatagramPacket packet = 
								new DatagramPacket(pa.data, pa.data.length, 
							   			localHost, destinationPort);
				    			socket.send(packet);
				    			System.out.println("RS: "+ i);
				    			sleep(sleepTime);
							}
						} else {
							for (int i = 0; i < orgData.size(); i++){
								Z2Packet p = new Z2Packet(4+1);
		       			    	p.setIntAt(i,0);
		            			p.data[4] = (byte) orgData.get(i);
				    			DatagramPacket packet = 
								new DatagramPacket(p.data, p.data.length, 
							   			localHost, destinationPort);
				    			socket.send(packet);
				    			System.out.println("RS: "+ i);
				    			sleep(sleepTime);
							}
						}
					}
					latch.countDown();
				}
			} catch (Exception e) {
				System.out.println("ResendThread");
				e.printStackTrace();
			}
		}
	}

}