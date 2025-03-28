// an interface type defines a set of method signatures.

// Yes! You've understood the core concept perfectly! Let me validate your understanding point by point:

// Interface as a Contract: Exactly right! Just like the socket defines "you must be 3-pin and work at 220V" to use it, a Go interface says "you must have these exact methods" to satisfy me.

// Real-world Analogy: Your TV/oven/microwave examples are spot on. In Go terms:

// The socket = interface

// TV/oven = concrete types (usually structs)

// Plug shape = method signatures

// Automatic Satisfaction: 100% correct! When you define a struct with methods that match the interface's requirements, it "plugs into" that interface automatically - no explicit declaration needed.

// One small clarification about terminology:

// The "signature" specifically refers to the method's name, parameters, and return types

// The "interface" is the complete set of required method signatures

package main

import "fmt"

 type Notifier interface {
	Notify(message string) // all notifiers must have this method. to use 
							// this interface
 }

 // now create concrete types.

 type EmailNotifier struct {
	Recipient string
 }

 // now use that interface like this.

 func(e EmailNotifier) Notify(message string) {
	fmt.Printf("Sending Email to %s: `%s`\n,",e.Recipient, message)
 }

 // another type

 type smsNotifier struct {
	PhoneNumber string
 }

 func (s smsNotifier) Notify(message string) {
	fmt.Printf("Sending SMS to %s : `%s`\n ", s.PhoneNumber,message)
 }

 // now we will write the function that will use that contract

 func sendNotification( n Notifier, message string) {
	fmt.Println("Dispatching notification -> ")
	n.Notify(message) // here we are calling the Notify method from the interface.

 }

 // now lets say we have to add another method for sending notification.

 // now we can just add it. like this

 type pushNotifier struct {
	DeviceId string
 }

 func(p pushNotifier) Notify(message string) {
	fmt.Printf("Sending push notification to %s : `%s`\n", p.DeviceId,message);
 }






func main(){

	// now lets put all together.

	email := EmailNotifier{Recipient: "indiagaurav@gmail.com"}
	sms := smsNotifier{PhoneNumber: "+91229292929"}
	push := pushNotifier{DeviceId: "device-abc-111"}

	messageToSend := "Your order has shipped!"

	

	fmt.Println("---this is the interface way---")

	// we can pass any notifier here to the send notification method.

	sendNotification(email,messageToSend)
	sendNotification(sms, messageToSend)
	sendNotification(push,messageToSend) 

	

}







