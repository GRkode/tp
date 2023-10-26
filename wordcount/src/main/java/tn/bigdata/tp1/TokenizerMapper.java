package tn.bigdata.tp1;

import org.apache.hadoop.io.DoubleWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Mapper;
import java.io.IOException;

public class TokenizerMapper extends Mapper<Object, Text, Text, DoubleWritable> {
    private Text store = new Text();
    private DoubleWritable cost = new DoubleWritable();

    public void map(Object key, Text value, Context context) throws IOException, InterruptedException {
        String[] fields = value.toString().split("\t");

        if (fields.length >= 6) {
            String storeName = fields[2];
            double saleCost = Double.parseDouble(fields[4]);

            store.set(storeName);
            cost.set(saleCost);

            context.write(store, cost);
        }
    }
}
