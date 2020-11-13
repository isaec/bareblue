//import the javafx stuff
import javafx.application.Application;
import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.StackPane;
import javafx.stage.Stage;

public class Main extends Application {


    Button button;

    public static void main(String[] args){
        launch(args);
    }

    @Override
    public void start(Stage primaryStage) throws Exception{
        primaryStage.setTitle("bareblue");

        button = new Button("I hope the ram use isnt too much..."); //new button


        StackPane layout = new StackPane();
        layout.getChildren().add(button);

        Scene scene = new Scene(layout, 300,300);
        primaryStage.setScene(scene);
        primaryStage.show();

        System.out.printf("%.3fGiB", Runtime.getRuntime().totalMemory() / (1024.0 * 1024.0 * 1024.0));}
}