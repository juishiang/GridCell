close all
clear all

%%%%%%%%parameter%%
space_size = 10*2;
cycle = 5;
k  =  space_size / cycle;
%%% matrix
T_d = (1/k)*[[1 -(1/sqrt(3))];[0 2/sqrt(3)]];
T_d_i = k*[[1 0.5];[0 sqrt(3)/2]];
offset = T_d_i*[rand;rand];
p_dev = 0.4; 
h_mean = 0.7;
h_dev = 0.6; 

%%% test data
x_s = 0:0.1:10;
y_s = 0:0.1:10;
fir_rate = zeros(length(x_s),length(y_s));
h_val = zeros(2*cycle+2,2*cycle+2);
for x1 = 1:1:length(x_s) 
    for y1 = 1:1:length(y_s)
        x = x_s(x1);
        y = y_s(y1);
        site = [x;y]-offset;
        p_idx = T_d * site;
        p_idx_start = floor(p_idx);
        p_site = p_idx_start;
        min = norm(site - (T_d_i*p_idx_start));
        for i = 0:1 
            for j = 0:1
                cur_s = norm(site - (T_d_i*(p_idx_start + [i;j])));
                if cur_s < min 
                    min = cur_s;
                    p_site = p_idx_start + [i;j];
                end
            end
        end
        p_site = p_site + [cycle+2;cycle+2];
        if h_val(p_site(1),p_site(2)) == 0
            h_val(p_site(1),p_site(2)) = h_mean + h_dev*randn;
            if h_val(p_site(1),p_site(2)) >1 
                h_val(p_site(1),p_site(2)) =1;
            elseif h_val(p_site(1),p_site(2)) <0;
                h_val(p_site(1),p_site(2)) = 0;
            end
        end
        fir_rate(y1,x1) = h_val(p_site(1),p_site(2))*exp(-(min^2)/(2*(p_dev)^2));%
    end
end
h = pcolor(x_s,y_s,fir_rate);
set(h, 'EdgeColor', 'none');