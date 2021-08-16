
DROP PROCEDURE IF EXISTS analyse_cost_month;
DELIMITER //
-- description 支出统计
-- author lujiang
-- created_at 2021-08-01
create procedure analyse_cost_month(in ym int)
begin
    DECLARE date_start date;
    DECLARE date_end date;


    set date_start = date_format(concat(ym,'01'),'%Y-%m-%d');
    set date_end = last_day(date_start);

    SET SQL_SAFE_UPDATES = 0;

    -- 按【user_id, ym, category_id】 进行统计
    delete from wealth_cost_statistics_months where `year_month` = ym;

    insert into wealth_cost_statistics_months(`year`, `month`, user_id,`year_month`,category_id,total)
    select 
        date_format(occur_date,'%Y'),
        date_format(occur_date,'%m'),
        user_id,
        date_format(occur_date,'%Y%m') ym, 
        category_id,
        sum(amount) total
    from wealth_cost_details
    where occur_date between date_start and date_end
    group by date_format(occur_date,'%Y'),user_id,date_format(occur_date,'%Y%m'), category_id;

end
//
DELIMITER ;
